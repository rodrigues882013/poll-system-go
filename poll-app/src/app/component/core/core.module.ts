import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TopbarComponent} from './topbar/topbar.component';
import {RouterModule} from '@angular/router';
import {TextCaptchaComponent} from './captcha/text-captcha/text-captcha.component';
import {FormsModule} from '@angular/forms';
import {ResultChartComponent} from './result-chart/result-chart.component';
import {ErrorPageComponent} from './error-page/error-page.component';

const components = [
  TopbarComponent,
  TextCaptchaComponent
];

@NgModule({
  imports: [CommonModule, RouterModule, FormsModule],
  exports: [components, ResultChartComponent, ErrorPageComponent],
  declarations: [components, ResultChartComponent, ErrorPageComponent]
})
export class CoreModule {}
