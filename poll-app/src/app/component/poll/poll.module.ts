import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {CommonModule} from '@angular/common';
import {PollPresentationComponent} from './poll-presentation/poll-presentation.component';
import {NominateCardComponent} from './nominate/nominate-card/nominate-card.component';
import {CoreModule} from '../core/core.module';
import {FormsModule} from '@angular/forms';


const routes: Routes = [
  {
    path: ':id',
    component: PollPresentationComponent
  }
];

@NgModule({
  imports: [
    RouterModule.forChild(routes),
    CommonModule,
    CoreModule,
    FormsModule
  ],
  exports: [
    RouterModule,
    PollPresentationComponent
  ],
  declarations: [
    PollPresentationComponent,
    NominateCardComponent,
  ]
})
export class PollModule {}
