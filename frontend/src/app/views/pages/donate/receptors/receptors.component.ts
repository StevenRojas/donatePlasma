import {Component, OnInit, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {RecipientService} from '../../../../core/donate/services/recipient.service';
import {RecipientModel} from '../../../../core/donate/models/recipient.model';
import {bloodTypes, cities} from '../../../../../environments/environment';
import {Observable} from 'rxjs';
import {MatSelectChange} from '@angular/material/select';
import {map} from 'rxjs/operators';
import {ActivatedRoute, Router} from '@angular/router';
import Swal, {SweetAlertOptions} from 'sweetalert2';
import {SwalComponent} from '@sweetalert2/ngx-sweetalert2';
import {DonorService} from '../../../../core/donate/services/donor.service';
import {Meta, Title} from "@angular/platform-browser";
import {MatPaginator} from "@angular/material/paginator";

@Component({
  selector: 'kt-receptors',
  templateUrl: './receptors.component.html',
  styleUrls: ['./receptors.component.scss']
})
export class ReceptorsComponent implements OnInit {
  @ViewChild('coolModal', {static: false}) private coolModal: SwalComponent;
  public coolModalOption: SweetAlertOptions;

  @ViewChild('failModal', {static: false}) private failModal: SwalComponent;
  public failModalOption: SweetAlertOptions;
  @ViewChild('listPaginator', {static: false}) public  listPaginator: MatPaginator;

  public formGroup: FormGroup;
  public removeFormGroup: FormGroup;
  public list: Observable<RecipientModel[]>;
  public bloodTypes = bloodTypes;
  public cities = cities;
  public loading: boolean;
  public page = 1;
  public size = 30;
  // Filters vars
  public query = '';
  public city = 0;
  public bloodType = 0;
  public total = 0;

  public donors = [];

  bloodTypeSelected = '';
  constructor(
    private router: Router,
    private fb: FormBuilder,
    private route: ActivatedRoute,
    private activatedRoute: ActivatedRoute,
    private recipientService: RecipientService,
    private donorService: DonorService,
    private title: Title,
    private meta: Meta
  ) {
    this.initRegisterFormGroup();
    this.title.setTitle('Dona tu Plasma - Lista de Receptores');
    this.meta.updateTag({ name: 'charset', content: 'UTF-8' });
    this.meta.updateTag({ name: 'description', content: '¿Quieres donar plasma o sabes quién podría? Recibe información sobre receptores de plasma para combatir el CODIV-19' });
    this.meta.updateTag({ name: 'robots', content: 'index, follow' });
    this.meta.updateTag({ property: 'og:url', content: 'https://donatuplasma.org/receptores' });
    this.meta.updateTag({ property: 'og:title', content: 'Dona tu Plasma - Lista de Receptores' });
    this.meta.updateTag({ property: 'og:description', content: '¿Quieres donar plasma o sabes quién podría? Recibe información sobre receptores de plasma para combatir el CODIV-19' });
    this.meta.updateTag({ property: 'og:image', content: 'https://donatuplasma.org/assets/media/plasma/og-imagen-receptores.jpg' });
    this.meta.updateTag({ property: 'og:image:width', content: '1200' });
    this.meta.updateTag({ property: 'og:image:height', content: '627' });
    this.meta.updateTag({ property: 'og:type', content: 'article' });
    this.meta.updateTag({ property: 'og:locale', content: 'es_ES' });
  }

  ngOnInit(): void {
    this.coolModalOption = {
      title: '',
      showCloseButton: true,
      showConfirmButton: false
    };
    this.failModalOption = {
      title: 'Eliminar Receptor',
      type: 'warning',
      showCloseButton: true,
      showCancelButton: true,
      showConfirmButton: true,
      confirmButtonText: 'Aceptar',
      cancelButtonText: 'Cancelar'
    };

    this.route.params.subscribe(params => {
      this.activatedRoute.queryParams.subscribe(queryParams => {
        if (queryParams.bt) {
          this.bloodType = queryParams.bt;
          this.formGroup.patchValue({
            blood_type: this.bloodType
          });
        }
        if(queryParams.city) {
          this.city = queryParams.city;
          this.formGroup.patchValue({
            city: this.city
          });
        }

        this.getAll();
      });
    });
  }

  public deleteReceptor(receptorId: number) {
    this.failModal.fire().then((result) => {
      if (result.value) {
        const answer = !!this.removeFormGroup.controls.answer.value;
        this.recipientService.delete(
          receptorId, answer, this.removeFormGroup.controls.comment.value).subscribe(
          (response) => {
            this.removeFormGroup.reset();
            Swal.fire({
              title: 'Gracias por su ayuda',
              type: 'success',
              timer: 3000
            }).then(() => {
              window.location.reload();
            });
          },
          (error) => this.removeFormGroup.reset()
        );
      }
    });
  }

  public showDonors(bloodType: any) {
    this.donorService.getDonorsByBloodType(bloodType).subscribe(
      (list) => {
        this.donors = list;
        this.bloodTypeSelected = bloodType;
        this.coolModal.fire().then((result) => {
          if (result.value) {
          }
        });
      }
    );
  }


  public initRegisterFormGroup() {
    this.formGroup = this.fb.group(
      {
        blood_type: ['0', Validators.compose([])],
        city: ['0', Validators.compose([])],
        query: ['', Validators.compose([])]
      }
    );
    this.removeFormGroup = this.fb.group({
      answer: ['', Validators.compose([])],
      comment: ['', Validators.compose([])],
    });
  }
  setQuery(text) {
    this.query = text;
    this.getAll();
  }

  setCity(optionSelected: MatSelectChange) {
    // tslint:disable-next-line:radix
    this.city = parseInt(optionSelected.value);
    this.getAll();
  }

  setBloodType(optionSelected: MatSelectChange) {
    // tslint:disable-next-line:radix
    this.bloodType = parseInt(optionSelected.value);
    this.getAll();
  }

  private getAll() {
    this.list = this.recipientService.search(this.page, this.size, this.city, this.bloodType, this.query)
        .pipe(
            map( result => {
              this.total = result.total_records;
              return result.recipients;
            })
        );
  }

  updatePagination(pagination: any) {
    this.size = pagination.pageSize;
    this.page = pagination.pageIndex + 1;
    this.getAll();
  }

  showImage(photoPath: string) {
    if (photoPath || photoPath !== '') {
      Swal.fire({
        imageUrl: `https://donatuplasma.org/assets/media/images/${photoPath}`,
      })
    } else {
      Swal.fire({
        text: 'No hay imágen',
      })
    }

  }
}
