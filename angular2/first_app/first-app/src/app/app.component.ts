import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styles: [`
    h1 {
      color: red;
    }
  `]
})
export class AppComponent {
  title = 'app works!';
  delete = false;
  test = 'Initial Value';
  boundValue = 1000;
}
