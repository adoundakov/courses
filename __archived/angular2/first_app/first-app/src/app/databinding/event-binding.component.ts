import { Component, OnInit, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-event-binding',
  template: `
    <button (click)='onClicked()'>Click Me!</button>
  `,
  styles: []
})

export class EventBindingComponent {
  @Output() clicked = new EventEmitter<string>();

  // onClicked () {
  //   // alert('It Worked!');
  //   console.log('It Worked!');
  // }

  onClicked () {
    this.clicked.emit('It works! Custom!!');
  }
}
