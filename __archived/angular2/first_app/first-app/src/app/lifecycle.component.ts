import { Component, OnInit, DoCheck, AfterContentInit, AfterContentChecked,
         AfterViewInit, AfterViewChecked, OnChanges, OnDestroy, Input, ViewChild,
         ContentChild
       } from '@angular/core';

@Component({
  selector: 'app-lifecycle',
  template: `
    <h1>
      Lifecycle Methods
    </h1>
    <ng-content></ng-content>
    <hr>
    <p #boundParagraph>Bindable = {{bindable}}</p>
    <p>Bound Paragraph Value = {{boundParagraph.textContent}}</p>
  `,
  styles: []
})
export class LifecycleComponent implements OnChanges, OnInit, DoCheck,
  AfterContentInit, AfterContentChecked, AfterViewInit, AfterViewChecked,
  OnDestroy {

  @Input() bindable = 1000;

  @ViewChild('boundParagraph') boundParagraph: HTMLElement;

  @ContentChild('boundContent') boundContent: HTMLElement;

  constructor() { }

  ngOnInit() {
    this.log('ngOnInit');
  }

  ngOnChanges() {
    this.log('ngOnChanges');
  }

  ngDoCheck() {
    this.log('ngDoCheck');
  }

  ngAfterContentInit() {
    this.log('ngAfterContentInit');
  }

  ngAfterContentChecked() {
    this.log('ngAfterContentChecked');
    console.log(this.boundContent);
  }

  ngAfterViewInit() {
    this.log('ngAfterViewInit');
  }

  ngAfterViewChecked() {
    this.log('ngAfterViewChecked');
    console.log(this.boundParagraph);
    // cant use this.log() above because this.log is expecting a string,
    // not an HTML element
  }

  ngOnDestroy() {
    this.log('ngOnDestroy');
  }

  private log(hook: string) {
    console.log(hook);
  }
}
