import {Injectable} from '@angular/core';
import {HttpClient, HttpParams} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../../../environments/environment';
import {map, retry} from 'rxjs/operators';
import {DonorModel} from '../models/donor.model';

@Injectable({
  providedIn: 'root'
})
export class DonorService {
  private tempRefObj = new DonorModel();

  constructor(
    private http: HttpClient
  ) {
  }


  public search(page: number, size: number, cityId: number, bloodType: number, query: string): Observable<any> {
    let params = new HttpParams();
    if (query && query !== '') {
      params = params.set('q', query);
    }
    if (size) {
      params = params.set('size', size.toString());
    }
    if (page) {
      params = params.set('page', page.toString());
    }
    if (cityId && cityId !== 0) {
      params = params.set('city', cityId.toString());
    }
    if (bloodType && bloodType !== 0) {
      params = params.set('compatible_with', bloodType.toString());
    }
    return this.http.get(environment.api_url + `donor/public`, {params})
      .pipe(
        retry(2),
        map((response: any) => response)
      );
  }


  public getDonorsByBloodType(bloodTypeId: number): Observable<any> {
    return this.http.get(environment.api_url_match + `donors/${bloodTypeId}`)
      .pipe(
        retry(2),
        map( (response: any) => response.donors )
      );
  }

  public getAllPublic(): Observable<DonorModel[]> {
    return this.http.get<DonorModel[]>(environment.api_url + 'donor/public')
      .pipe(
        retry(2),
        map((response: any) => {
          const list = [];
          Object.keys(response.donors).forEach((item, key) => {
            list.push(Object.assign(new DonorModel(), response.donors[key]))
          });
          return list;
        })
      );
  }

  public getCanReceiveFrom(bloodType: number): Observable<[]> {
    return this.http.get<[]>(environment.api_url_match + `can-donate-to/${bloodType}`)
      .pipe(
        retry(2),
        map((response: any) => {
          return response.compatible_types
        })
      );
  }

  public get(donorId: number): Observable<DonorModel> {
    return this.http.get<DonorModel>(environment.api_url + 'donor/' + donorId)
      .pipe(
        retry(2),
        map((response: any) => Object.assign(new DonorModel(), response[this.tempRefObj.getPrefix()]))
      );
  }

  public post(donor: any): Observable<DonorModel> {
    const body = {};
    body[this.tempRefObj.getPrefix()] = donor;
    return this.http.post<DonorModel>(environment.api_url + 'donor', body)
      .pipe(
        map((response: any) => Object.assign(new DonorModel(), response[this.tempRefObj.getPrefix()]))
      );
  }

  public put(donor: DonorModel): Observable<DonorModel> {
    const body = {};
    body[this.tempRefObj.getPrefix()] = donor;
    return this.http.put<DonorModel>(environment.api_url + 'donor/' + donor.id, body)
      .pipe(
        map((response: any) => Object.assign(new DonorModel(), response[this.tempRefObj.getPrefix()]))
      );
  }

  public delete(donorId: number): Observable<any> {
    return this.http.delete<any>(environment.api_url + `donor/${donorId}`)
      .pipe(
        retry(2),
        map(response => donorId)
      );
  }

  public getManagement() {
    return this.http.get<[]>(environment.api_url + 'donor')
      .pipe(
        retry(2),
        map((response: any) => response.donors)
      );
  }

  public verify(donorId: number, isVerify: boolean): Observable<boolean> {
    return this.http.patch<boolean>(environment.api_url + `donor/${donorId}/verify`,
      {verified: isVerify})
      .pipe(
        retry(2),
        map((response: any) => response.ok)
      );
  }

  public setPublic(donorId: number, isPublic: boolean): Observable<boolean> {
    return this.http.patch<boolean>(environment.api_url + `donor/${donorId}/public`,
      {public: isPublic})
      .pipe(
        retry(2),
        map((response: any) => response.ok)
      );
  }


  public setActivate(donorId: number, isActivate: boolean): Observable<boolean> {
    return this.http.patch<boolean>(environment.api_url + `donor/${donorId}/activate`,
      {activate: isActivate})
      .pipe(
        retry(2),
        map((response: any) => response.ok)
      );
  }

}
