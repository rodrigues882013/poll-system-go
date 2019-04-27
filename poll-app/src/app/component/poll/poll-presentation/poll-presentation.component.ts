import {Component, OnInit} from '@angular/core';
import {PollService} from '../poll.service';
import {ActivatedRoute} from '@angular/router';
import {Poll} from '../../core/models/Poll';
import {Nominate} from '../../core/models/Nominate';
import {Vote} from '../../core/models/Vote';
import {VoteConfirmationComponent} from '../vote/vote-confirmation/vote-confirmation.component';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-poll-presentation',
  templateUrl: './poll-presentation.component.html',
  styleUrls: ['./poll-presentation.component.scss']
})
export class PollPresentationComponent implements OnInit {
  poll: Poll;
  selected: Nominate;
  chosen: boolean[];
  vote: Vote = new Vote();
  success = '';
  data = {result: []};
  hiddenElement = false;

  constructor(private pollService: PollService,
              private route: ActivatedRoute,
              private modalService: NgbModal) {
    this.poll = new Poll();
  }

  ngOnInit() {
    this.pollService
      .getPollById(+this.route.snapshot.paramMap.get('id'))
      .subscribe((data) => {

        if (data) {
          this.poll = data;
          this.chosen = this.poll.nominates.map(() => false);
        }

      });
  }

  changeSelection(nominate) {
    this.selected = nominate;
    this.chosen = this
      .chosen
      .map((_, i) => i === this.poll.nominates.findIndex(x  => x.id === nominate.id));
  }

  clearField() {
    this.vote.name = '';
    this.vote.email = '';
  }

  saveVote() {
    this.pollService
      .vote(this.vote.poll.id, this.vote)
      .subscribe(x => this.voteSuccessCallBack(x));
  }

  voteSuccessCallBack(res) {
    if (res.status === 202) {

      this.pollService
        .getPollResult(+this.route.snapshot.paramMap.get('id'))
        .subscribe((data) => {
          this.data.result = data.results.map(x => ['user: ' + x.nominate.id, x.votes]);
          this.hiddenElement = !this.hiddenElement;
          this.success = `<strong>Parabens!</strong> Seu voto no <strong>${this.selected.name}</strong> foi enviado com sucesso.`;
          this.chosen = this.chosen.map(() => false);
        });
    }
  }

  open() {
    this.vote.nominate = this.selected;
    this.vote.poll = this.poll;

    const modalRef = this.modalService.open(VoteConfirmationComponent);
    modalRef.componentInstance.vote = this.vote;

    modalRef
      .result
      .then(() => this.saveVote())
      .catch(() => console.log('error'))
      .finally(() => this.clearField());
  }
}
