import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-error-page-component',
  templateUrl: './error-page.component.html',
  styleUrls: ['./error-page.component.scss']
})
export class ErrorPageComponent implements OnInit{

  error: string;

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    const cod = +this.route.snapshot.paramMap.get('status');

    switch (cod) {
      case 403:
        this.error = 'Votação encerrada.';
        break;
      case 404:
        this.error = 'Não existe votação';
        break;
      default:
        this.error = 'Algo deu errado';
        break;
    }
  }



}
