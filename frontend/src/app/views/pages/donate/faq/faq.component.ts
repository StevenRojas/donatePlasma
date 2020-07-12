import {Component, OnInit, ViewChild} from '@angular/core';
import Swal, {SweetAlertOptions} from 'sweetalert2';
import {SwalComponent} from '@sweetalert2/ngx-sweetalert2';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { FaqService } from 'src/app/core/donate/services/faq.service';
import {Meta, Title} from '@angular/platform-browser';

@Component({
  selector: 'kt-faq',
  templateUrl: './faq.component.html',
  styleUrls: ['./faq.component.scss']
})
export class FaqComponent implements OnInit {
  public coolModalOption: SweetAlertOptions;
  public suggestModalOption: SweetAlertOptions;
  @ViewChild('coolModal', {static: false}) private coolModal: SwalComponent;
  @ViewChild('suggestModal', {static: false}) private suggestModal: SwalComponent;
  public coolFormGroup: FormGroup;
  public suggestFormGroup: FormGroup;
  constructor(
      private fb: FormBuilder,
      private faqService: FaqService,
      private title: Title,
      private meta: Meta
  ) {
    this.initRegisterFormGroup();
    this.title.setTitle('Dona tu Plasma - Infórmate');
    this.meta.updateTag({ name: 'charset', content: 'UTF-8' });
    this.meta.updateTag({ name: 'description', content: 'Infórmate para saber cómo ayudar a donar o recibir plasma y salvar vidas combatiendo el COVID-19.' });
    this.meta.updateTag({ name: 'robots', content: 'index, follow' });
    this.meta.updateTag({ property: 'og:url', content: 'https://donatuplasma.org/faq' });
    this.meta.updateTag({ property: 'og:title', content: 'Dona tu Plasma -  Infórmate' });
    this.meta.updateTag({ property: 'og:description', content: 'Infórmate para saber cómo ayudar a donar o recibir plasma y salvar vidas combatiendo el COVID-19.' });
    this.meta.updateTag({ property: 'og:image', content: 'https://donatuplasma.org/assets/media/plasma/og-imagen-faq.jpg' });
    this.meta.updateTag({ property: 'og:image:width', content: '1200' });
    this.meta.updateTag({ property: 'og:image:height', content: '627' });
    this.meta.updateTag({ property: 'og:type', content: 'article' });
    this.meta.updateTag({ property: 'og:locale', content: 'es_ES' });
  }

  ngOnInit(): void {
    this.coolModalOption = {
      title: 'Quiero ser voluntario!',
      type: 'success',
      showCancelButton: true,
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Quiero ayudar',
      focusCancel: true,
      showLoaderOnConfirm: true,
      preConfirm: () => this.sendCool()
    };

    this.suggestModalOption = {
      title: 'Mándanos tus sugerencias',
      type: 'success',
      showCancelButton: true,
      cancelButtonText: 'Cancelar',
      confirmButtonText: 'Enviar',
      focusCancel: true,
      showLoaderOnConfirm: true,
      preConfirm: () => this.sendSuggest()
    };
  }


  /**
   * Modal for Delete Obj
   * @param event
   */
  public openCoolModal(event) {
    this.coolModal.fire().then((result) => {
      if (result.value) {
        Swal.fire(
          'Gracias!',
          'Nos pondremos en contacto contigo',
          'success'
        )
      }
    });
  }

  /**
   * Modal for Delete Obj
   * @param event
   */
  public openSuggestModel(event) {
    this.suggestModal.fire().then((result) => {
      if (result.value) {
        Swal.fire(
          'Gracias!',
          'Muchas gracias por su sugerencia',
          'success'
        )
      }
    });
  }

  public initRegisterFormGroup() {
    this.coolFormGroup = this.fb.group(
        {
          alias: ['', Validators.compose([Validators.required, Validators.email])],
          comment: ['', Validators.compose([Validators.required, Validators.minLength(10)])]
        }
    );
    this.suggestFormGroup = this.fb.group(
        {
          alias: ['', Validators.compose([])],
          comment: ['', Validators.compose([Validators.required, Validators.minLength(10)])]
        }
    );

  }

  sendCool() {
    const controls = this.coolFormGroup.controls;
    // check form
    if (this.coolFormGroup.invalid) {
      Object.keys(controls).forEach(controlName =>
          controls[controlName].markAsTouched()
      );
      return false;
    }
    const postObj = this.coolFormGroup.getRawValue();
    const response = this.faqService.recruit(postObj.alias, postObj.comment);
    return new Promise((resolve, reject) => {
      response.subscribe(
        // tslint:disable-next-line:no-shadowed-variable
        (response) => {
          resolve();
        },
        error => reject()
      );
    });
  }

  sendSuggest() {
    const controls = this.suggestFormGroup.controls;
    // check form
    if (this.suggestFormGroup.invalid) {
      Object.keys(controls).forEach(controlName =>
          controls[controlName].markAsTouched()
      );
      return false;
    }
    const postObj = this.suggestFormGroup.getRawValue();
    const response = this.faqService.comments(postObj.alias, postObj.comment);
    return new Promise((resolve, reject) => {
      response.subscribe(
        // tslint:disable-next-line:no-shadowed-variable
        (response) => {
          resolve();
        },
        error => reject()
      );
    });

  }
  public isControlHasErrorCool(controlName: string, validationType: string): boolean {
    const control = this.coolFormGroup.controls[controlName];
    if (!control) {
      return false;
    }
    return control.hasError(validationType) && (control.dirty || control.touched);
  }

  public isControlHasErrorSuggest(controlName: string, validationType: string): boolean {
    const control = this.suggestFormGroup.controls[controlName];
    if (!control) {
      return false;
    }
    return control.hasError(validationType) && (control.dirty || control.touched);
  }
}
