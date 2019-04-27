import {Component} from '@angular/core';
import {NgbActiveModal} from '@ng-bootstrap/ng-bootstrap';
import {PollService} from '../../poll.service';
import {ActivatedRoute} from '@angular/router';
import {Vote} from '../../../core/models/Vote';

@Component({
  selector: 'app-vote-confirmation',
  templateUrl: './vote-confirmation.component.html',
  styleUrls: ['./vote-confirmation.component.scss']
})
export class VoteConfirmationComponent {
  vote: Vote;
  canBeVote = false;

  constructor(public activeModal: NgbActiveModal) {}

  submitVote() {
    this.activeModal.close('Success');
  }

  checkAnswer(isCorrect: boolean) {
    this.canBeVote = isCorrect;
  }

  isWrongEmail() {
    return this.vote.email.match('[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+') === null;
  }
}
