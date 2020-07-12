import {Component, OnInit, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {bloodTypes, cities, environment} from '../../../../../environments/environment';
import {SwalComponent} from '@sweetalert2/ngx-sweetalert2';
import {SweetAlertOptions} from 'sweetalert2';
import {RecipientService} from '../../../../core/donate/services/recipient.service';
import {RecipientModel} from '../../../../core/donate/models/recipient.model';
import {MatSelectChange} from '@angular/material/select';
import {Meta, Title} from "@angular/platform-browser";
import {Router} from "@angular/router";

@Component({
  selector: 'kt-receptor',
  templateUrl: './receptor.component.html',
  styleUrls: ['./receptor.component.scss']
})
export class ReceptorComponent implements OnInit {
  public pondFiles = [];
  public lastFileAdd = '';
  public pondOptions = this.optionsFile();
  @ViewChild('myPond', { static: false }) myPond: any;
  public formGroup: FormGroup;
  public bloodTypes = bloodTypes;
  public cities = cities;
  public loading: boolean;
  public donors: any[];
  public bloodTypeSelected: number;

  @ViewChild('coolModal', {static: false}) private coolModal: SwalComponent;
  public coolModalOption: SweetAlertOptions;

  @ViewChild('failModal', {static: false}) private failModal: SwalComponent;
  public failModalOption: SweetAlertOptions;

  @ViewChild('tcModal', {static: false}) private tcModal: SwalComponent;
  public tcModalOption: SweetAlertOptions;
  constructor(
      private router: Router,
      private fb: FormBuilder,
      private recipientService: RecipientService,
      private title: Title,
      private meta: Meta
  ) {
    this.initRegisterFormGroup();
    this.title.setTitle('Dona tu Plasma - Regístrate si necesitas plasma');
    this.meta.updateTag({ name: 'charset', content: 'UTF-8' });
    this.meta.updateTag({ name: 'description', content: '¿Necesitas recibir plasma? Llena tu información aquí.' });
    this.meta.updateTag({ name: 'robots', content: 'index, follow' });
    this.meta.updateTag({ property: 'og:url', content: 'https://donatuplasma.org/recibe' });
    this.meta.updateTag({ property: 'og:title', content: 'Recibe Plasma' });
    this.meta.updateTag({ property: 'og:description', content: '¿Necesitas recibir plasma? Llena tu información aquí.' });
    this.meta.updateTag({ property: 'og:image', content: 'https://donatuplasma.org/assets/media/plasma/og-imagen-recibe.jpg' });
    this.meta.updateTag({ property: 'og:image:width', content: '1200' });
    this.meta.updateTag({ property: 'og:image:height', content: '627' });
    this.meta.updateTag({ property: 'og:type', content: 'article' });
    this.meta.updateTag({ property: 'og:locale', content: 'es_ES' });
  }

  ngOnInit(): void {
    this.coolModalOption = {
      title: 'No perdamos la esperanza!',
      type: 'success',
      showCloseButton: false,
      showConfirmButton: false,
      timer: 5000
    };

    this.failModalOption = {
      title: ':( Algo salio mal',
      type: 'error',
      showCloseButton: false,
      showConfirmButton: false,
      timer: 10000
    };

    this.tcModalOption = {
      title: 'Términos y condiciones',
      type: 'info',
      showCloseButton: false,
      showConfirmButton: true
    };
  }

  public initRegisterFormGroup() {
    this.formGroup = this.fb.group(
        {
          blood_type_id: ['', Validators.compose([Validators.required])],
          name: ['', Validators.compose([])],
          cell_phones: ['', Validators.compose([Validators.required])],
          email: ['', Validators.compose([Validators.email])],
          city_id: ['', Validators.compose([])],
          public: [true, Validators.compose([])],
          tandc: ['', Validators.compose([Validators.required, Validators.requiredTrue])]
        }
    );
  }

  public isControlHasError(controlName: string, validationType: string): boolean {
    const control = this.formGroup.controls[controlName];
    if (!control) {
      return false;
    }
    return control.hasError(validationType) && (control.dirty || control.touched);
  }

  submit() {
    const controls = this.formGroup.controls;
    // check form
    if (this.formGroup.invalid) {
      Object.keys(controls).forEach(controlName =>
          controls[controlName].markAsTouched()
      );
      return;
    }
    this.loading = true;
    const postObj = this.formGroup.getRawValue();
    postObj.photo_path = this.lastFileAdd;
    Object.keys(postObj).forEach(controlName => {
      if (['blood_type_id', 'city_id'].indexOf(controlName) !== -1) {
        if(postObj[controlName] === '') {
          delete postObj[controlName];
        } else {
          postObj[controlName] = parseInt(postObj[controlName], 10);
        }
      } else {
        postObj[controlName] = postObj[controlName];
      }
    });
    this.recipientService.post(postObj).subscribe(
        (donor: RecipientModel) => {
          this.coolModal.fire().then((result) => {
            if (result.value) {
            }
          });
          this.loading = false;
          this.formGroup.reset();
          setTimeout(() => {
            this.router.navigate(['/receptores']);
          }, 1000);
        },
        error => {
          this.failModal.fire().then((result) => {
            if (result.value) {
            }
          });
        }
    );
  }


  public getDonors(option: MatSelectChange) {
    this.bloodTypeSelected = option.value;
    this.recipientService.getCanReceiveFrom(option.value).subscribe(
        (possibleDonors) => {
          this.donors = possibleDonors;
        }
    );
  }

  /**
   * Configuration uploader
   *  @return configuration filepond
   */
  public optionsFile() {
    const token = 123;
    return {
      class: 'poi-file_uploader',
      multiple: false,
      labelIdle: 'Arrastre y suelte el archivo aquí o puede  <a class="link"> buscarlos </a>',
      acceptedFileTypes: 'image/*',
      instantUpload: true,
      maxFileSize: '5MB',
      allowRevert: false,
      server: {
        headers: {
          Authorization: `Bearer ${token}`,
          token
        },
        // ADD endpoit for upload photo
        process: (fieldName, file, metadata, load, error, progress, abort) => {
          const formData = new FormData();
          formData.append('file_uploader', file, file.name);
          const request = new XMLHttpRequest();
          request.open('POST', environment.api_url_simple + '/api/register/uploader');
          // request.setRequestHeader('Authorization', `Bearer ${token}`);
          // request.setRequestHeader('token', String(token));
          request.upload.onprogress = (e) => {
            progress(e.lengthComputable, e.loaded, e.total);
          };
          request.onload = () => {
            this.requestOnLoad(load, request, error);
          };
          request.send(formData);
          return {
            abort: () => {
              request.abort();
              abort();
            }
          };
        },
        revert: false
      },
      onremovefile: (error, file) => {
        this.lastFileAdd = '';
        // Use this method for remove file after upload, remove from server
      }
    };
  }

  public openTCModal() {
    this.tcModal.fire().then((result) => {
      if (result.value) {

      }
    });
  }
  /**
   * Remove File after upload document
   * @param event filepond event
   */
  public processFile(event: any) {
  }

  private requestOnLoad(load: any, request: any, error: any) {
    if (request.status >= 200 && request.status < 300) {
      const response = JSON.parse(request.responseText);
      this.lastFileAdd = response.filename;
      load(request.responseText);
    } else {
      error('El servicio no esta Disponible en este momento');
    }
  }
}
