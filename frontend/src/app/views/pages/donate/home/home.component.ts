import { Component, OnInit } from '@angular/core';
import { Meta, Title } from '@angular/platform-browser';

@Component({
  selector: 'kt-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  constructor(
    private title: Title,
    private meta: Meta
  ) {
    this.title.setTitle('Dona tu Plasma y Salva Vidas');
    this.meta.updateTag({ name: 'charset', content: 'UTF-8' });
    this.meta.updateTag({ name: 'description', content: 'Recibe información sobre donantes y receptores de plasma para combatir el CODIV-19' });
    this.meta.updateTag({ name: 'robots', content: 'index, follow' });
    this.meta.updateTag({ property: 'og:url', content: 'https://donatuplasma.org' });
    this.meta.updateTag({ property: 'og:title', content: 'Dona tu Plasma y Salva Vidas' });
    this.meta.updateTag({ property: 'og:description', content: 'Recibe información sobre donantes y receptores de plasma para combatir el CODIV-19' });
    this.meta.updateTag({ property: 'og:image', content: 'https://donatuplasma.org/assets/media/plasma/og-imagen-home.jpg' });
    this.meta.updateTag({ property: 'og:image:width', content: '1200' });
    this.meta.updateTag({ property: 'og:image:height', content: '627' });
    this.meta.updateTag({ property: 'og:type', content: 'website' });
    this.meta.updateTag({ property: 'og:locale', content: 'es_ES' });
  }

  ngOnInit(): void {
  }

}
