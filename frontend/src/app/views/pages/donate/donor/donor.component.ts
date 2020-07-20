import {Component, OnInit, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {bloodTypes, cities} from "../../../../../environments/environment";
import {SwalComponent} from "@sweetalert2/ngx-sweetalert2";
import {SweetAlertOptions} from 'sweetalert2';
import {DonorService} from "../../../../core/donate/services/donor.service";
import {DonorModel} from "../../../../core/donate/models/donor.model";
import {Router} from "@angular/router";
import {MatSelectChange} from "@angular/material/select";
import {MatCheckboxChange} from "@angular/material/checkbox";
import {Meta, Title} from "@angular/platform-browser";

@Component({
  selector: 'kt-donor',
  templateUrl: './donor.component.html',
  styleUrls: ['./donor.component.scss']
})
export class DonorComponent implements OnInit {
  @ViewChild('coolModal', {static: false}) private coolModal: SwalComponent;
  public coolModalOption: SweetAlertOptions;

  @ViewChild('failModal', {static: false}) private failModal: SwalComponent;
  public failModalOption: SweetAlertOptions;

  @ViewChild('tcModal', {static: false}) private tcModal: SwalComponent;
  public tcModalOption: SweetAlertOptions;

  public formGroup: FormGroup;
  public bloodTypes = bloodTypes;
  public cities = cities;
  public loading: boolean;
  public disabledAcceptTandC = true;
  public receptors: any[];
  public bloodTypeSelected: number;
  public quantityDonations = [
    { value: '1', display: 'Una sola vez'},
    { value: '1-3', display: 'De 1 a 3 veces'},
    { value: '3-6', display: 'De 3 a 6 veces'},
    { value: '6+', display: 'Más de 6 veces'},
    { value: 'no', display: 'No lo sé'}
  ];

  constructor(
    private router: Router,
    private fb: FormBuilder,
    private donorService: DonorService,
    private title: Title,
    private meta: Meta
  ) {
    this.initRegisterFormGroup();
    this.title.setTitle('Dona tu Plasma - Regístrate');
    this.meta.updateTag({ name: 'charset', content: 'UTF-8' });
    this.meta.updateTag({ name: 'description', content: '¿Quieres donar plasma? Llena tu información aquí para donar plasma.' });
    this.meta.updateTag({ name: 'robots', content: 'index, follow' });
    this.meta.updateTag({ property: 'og:url', content: 'https://donatuplasma.org/dona' });
    this.meta.updateTag({ property: 'og:title', content: 'Dona tu Plasma -  Llena tu información' });
    this.meta.updateTag({ property: 'og:description', content: '¿Quieres donar plasma? Llena tu información aquí.' });
    this.meta.updateTag({ property: 'og:image', content: 'https://donatuplasma.org/assets/media/plasma/og-imagen-dona.jpg' });
    this.meta.updateTag({ property: 'og:image:width', content: '1200' });
    this.meta.updateTag({ property: 'og:image:height', content: '627' });
    this.meta.updateTag({ property: 'og:type', content: 'article' });
    this.meta.updateTag({ property: 'og:locale', content: 'es_ES' });
  }

  ngOnInit(): void {
    this.coolModalOption = {
      title: 'Gracias!',
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
        cell: ['', Validators.compose([Validators.required])],
        email: ['', Validators.compose([Validators.email])],
        city_id: ['', Validators.compose([])],
        public: [true, Validators.compose([])],
        quantity_donations: ['', Validators.compose([Validators.required])],
        tandc: ['', Validators.compose([Validators.required, Validators.requiredTrue])]
      }
    );
  }

  public getReceptors(option: MatSelectChange) {
    this.bloodTypeSelected = option.value;
    this.donorService.getCanReceiveFrom(option.value).subscribe(
      (possibleDonors) => {
        this.receptors = possibleDonors;
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
    this.donorService.post(postObj).subscribe(
        (donor: DonorModel) => {
          this.coolModal.fire().then((result) => {
            if (result.value) {
            }
          });
          this.loading = false;
          this.formGroup.reset();
          const queryParams = {
            bt: postObj.blood_type_id,
          };
          if (postObj.city_id && postObj.city_id !== 0) {
            // @ts-ignore
            queryParams.city = postObj.city_id
          }
          this.router.navigate(['/donadores'], {queryParams});
        },
        error => {
          this.failModal.fire().then((result) => {
            if (result.value) {
            }
          });
        }
    );
  }

  checkCAndT($event: MatCheckboxChange) {
    this.disabledAcceptTandC = $event.checked;
  }

  public openTCModal() {
    this.tcModal.fire().then((result) => {
      if (result.value) {

      }
    });
  }
}
