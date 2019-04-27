import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ErrorPageComponent} from './component/core/error-page/error-page.component';


const routes: Routes = [
  {
    path: 'polls',
    loadChildren: './component/poll/poll.module#PollModule'
  },
  {
    path: 'error',
    component: ErrorPageComponent
  }
];

@NgModule({
  exports: [ RouterModule ],
  imports: [RouterModule.forRoot(routes)]
})


export class AppRoutingModule { }
