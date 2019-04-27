import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable, of} from 'rxjs';
import {catchError, tap} from 'rxjs/operators';
import {Captcha} from '../models/Captcha';

@Injectable({
  providedIn: 'root'
})
export class CaptchaService {
  // URL = 'http://0.0.0.0:5000';
  URL = 'http://localhost/captcha';

  constructor(private http: HttpClient) { }

  getCaptcha(): Observable<Captcha> {
    const url = `${this.URL}/mail.json`;

    return this.http
      .get<Captcha>(url)
      .pipe(
        tap(() => CaptchaService.log('fetched captcha')),
        catchError(this.handleError<Captcha>('getCaptcha')));
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      CaptchaService.log(`${operation} failed: ${error.message}`);
      return of(result as T);
    };
  }

  /** Log a HeroService message with the MessageService */
  private static log(message: string) {
    console.log(`HeroService: ${message}`);
  }

}
