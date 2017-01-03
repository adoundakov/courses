import { Component, OnInit } from '@angular/core';
import { Recipe } from '../recipe';
import { AppComponent } from '../../app.component';

@Component({
  selector: 'rb-recipe-list',
  templateUrl: './recipe-list.component.html'
})
export class RecipeListComponent implements OnInit {
  recipes: Recipe[] = [];
  recipe = new Recipe('Dummy', 'Dummy', 'https://cdn0.iconfinder.com/data/icons/kameleon-free-pack-rounded/110/Food-Dome-128.png');
  constructor() { }

  ngOnInit() {
  }

}
