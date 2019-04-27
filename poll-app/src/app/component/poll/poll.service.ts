import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable, of} from 'rxjs';
import {catchError, tap} from 'rxjs/operators';
import {PollResult} from '../core/models/PollResult';
import {Vote} from '../core/models/Vote';
import {Poll} from '../core/models/Poll';
import {NavigationExtras, Router} from '@angular/router';

@Injectable({providedIn: 'root'})
export class PollService {

  constructor(private http: HttpClient,
              private route: Router) { }
  URL = 'http://localhost/poll-api';
  // URL = 'http://www.mocky.io/v2/5cbce872320000341f80d96d';
  // result -> http://www.mocky.io/v2/5cbdbcdf2f0000740c16ce98

  private static log(message: string) {
    console.log(`PollService: ${message}`);
  }

  getPollResult(pollId: number): Observable<PollResult> {
    const url = `${this.URL}/polls/${pollId}/results?byNominate=true`;
    return this.http
      .get<PollResult>(url)
      .pipe(
        tap(() => PollService.log('fetch result')),
        catchError(this.handleError<PollResult>('getPollResult')));
  }

  getPollById(pollId: number): Observable<Poll> {
    const url = `${this.URL}/polls/${pollId}`;
    return this.http
      .get<Poll>(url)
      .pipe(
        tap(() => PollService.log('fetch poll')),
        catchError(this.handleError<Poll>('getPollById')));
  }

  vote(pollId: number, vote: Vote) {
    const url = `${this.URL}/polls/${pollId}/vote`;
    return this.http
      .post(url, vote, {observe: 'response'})
      .pipe(
        catchError(this.handleError<Vote>('vote')));
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      this.route.navigate([`/error`, {status: error.status}]);
      PollService.log(`${operation} failed: ${error.message}`);
      return of(result as T);
    };
  }
}
