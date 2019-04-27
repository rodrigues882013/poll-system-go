import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-nominate-card',
  templateUrl: './nominate-card.component.html',
  styleUrls: ['./nominate-card.component.scss']
})
export class NominateCardComponent implements OnInit {
  @Input() nominate;
  @Input() cardChosen;
  @Output() select = new EventEmitter();
  constructor() { }

  ngOnInit() {
  }

  changeSelection() {
    this.select.emit(this.nominate);
  }
}
