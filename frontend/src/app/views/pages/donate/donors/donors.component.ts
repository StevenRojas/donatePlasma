import {Component, OnChanges, OnInit, SimpleChanges, ViewChild} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import {RecipientService} from '../../../../core/donate/services/recipient.service';
import {bloodTypes, cities} from '../../../../../environments/environment';
import {ActivatedRoute, Router} from '@angular/router';
import Swal, {SweetAlertOptions} from 'sweetalert2';
import {SwalComponent} from '@sweetalert2/ngx-sweetalert2';
import {DonorService} from '../../../../core/donate/services/donor.service';
import {Meta, Title} from "@angular/platform-browser";
import {map} from "rxjs/operators";
import {Observable} from "rxjs";
import {DonorModel} from "../../../../core/donate/models/donor.model";
import {MatPaginator} from "@angular/material/paginator";

@Component({
  selector: 'kt-donors',
  templateUrl: './donors.component.html',
  styleUrls: ['./donors.component.scss']
})
export class DonorsComponent implements OnInit, OnChanges {

  @ViewChild('coolModal', {static: false}) private coolModal: SwalComponent;
  public coolModalOption: SweetAlertOptions;

  @ViewChild('failModal', {static: false}) private failModal: SwalComponent;

  @ViewChild('listPaginator', {static: false}) public  listPaginator: MatPaginator;
  public failModalOption: SweetAlertOptions;

  public formGroup: FormGroup;
  public removeFormGroup: FormGroup;
  public list: Observable<DonorModel[]>;
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
    this.title.setTitle('Dona tu Plasma - Lista de Donantes');
    this.meta.updateTag({ name: 'charset', content: 'UTF-8' });
    this.meta.updateTag({ name: 'description', content: '¿Necesitas recibir plasma? Aquí encontrarás información sobre posibles donantes de plasma para combatir el COVID-19.' });
    this.meta.updateTag({ name: 'robots', content: 'index, follow' });
    this.meta.updateTag({ property: 'og:url', content: 'https://donatuplasma.org/donadores' });
    this.meta.updateTag({ property: 'og:title', content: 'Dona tu Plasma -  Lista de Donantes' });
    this.meta.updateTag({ property: 'og:description', content: '¿Necesitas recibir plasma? Aquí encontrarás información sobre posibles donantes de plasma para combatir el COVID-19.' });
    this.meta.updateTag({ property: 'og:image', content: 'https://donatuplasma.org/assets/media/plasma/og-imagen-donadores.jpg' });
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
      title: 'Eliminar Donador',
      type: 'warning',
      showCloseButton: true,
      showCancelButton: true,
      showConfirmButton: true,
      confirmButtonText: 'Aceptar',
      cancelButtonText: 'Cancelar'
    };
    this.getAll();
  }

  ngOnChanges(changes: SimpleChanges): void {
  }

  public deleteReceptor(receptorId: number) {
    this.failModal.fire().then((result) => {
      if (result.value) {
        const answer = !!this.removeFormGroup.controls.answer.value;
        this.donorService.delete(
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


  public initRegisterFormGroup() {
    this.formGroup = this.fb.group(
      {
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


  private getAll() {
    this.list = this.donorService.search(this.page, this.size, this.city, this.bloodType, this.query)
        .pipe(
            map( result => {
              this.total = result.total_records;
              return result.donors;
            })
        );
  }

  updatePagination(pagination: any) {
    this.size = pagination.pageSize;
    this.page = pagination.pageIndex + 1;
    this.getAll();
  }

}
