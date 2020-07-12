import { Component, OnInit } from '@angular/core';
import { HtmlClassService } from '../../../theme/html-class.service';

@Component({
  selector: 'kt-footer-plasma',
  templateUrl: './footer-plasma.component.html',
  styleUrls: ['./footer-plasma.component.scss']
})
export class FooterPlasmaComponent implements OnInit {
// Public properties
  today: number = Date.now();
  footerClasses = '';
  footerContainerClasses = '';
  constructor(
      private uiClasses: HtmlClassService
  ) { }

  ngOnInit(): void {
    this.footerClasses = this.uiClasses.getClasses('footer', true).toString();
    this.footerContainerClasses = this.uiClasses.getClasses('footer_container', true).toString();
  }

}
