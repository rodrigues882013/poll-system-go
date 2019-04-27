import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {Md5} from 'ts-md5';
import {CaptchaService} from '../captcha.service';
import {Captcha} from '../../models/Captcha';

@Component({
  selector: 'app-text-captcha',
  templateUrl: './text-captcha.component.html',
  styleUrls: ['./text-captcha.component.scss']
})
export class TextCaptchaComponent implements OnInit {

  @Output() canBeVote = new EventEmitter<boolean>();

  questions = [];
  chosen: Captcha;
  answer = '';

  constructor(private captchaService: CaptchaService) {}

  ngOnInit() {
    this.getQuestion();
  }

  checkAnswer() {
    const isCorrect = this.chosen.a.some(x => x === Md5.hashAsciiStr(this.answer.toLowerCase()));

    if (isCorrect) {
      this.canBeVote.emit(isCorrect);
    } else {
      this.answer = '';
      this.getQuestion();
    }
  }

  getQuestion() {
    this.captchaService
      .getCaptcha()
      .subscribe((data) => {
        this.chosen = data;
      });

    this.chosen = this.questions[Math.floor(Math.random() * this.questions.length)];
  }

}
