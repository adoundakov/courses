import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-another',
  template: `
    <article>
      <ng-content></ng-content>
    </article>
  `,
  styles: [`
      article {
        border: 1px solid blue;
      }
    `]
})
export class AnotherComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
