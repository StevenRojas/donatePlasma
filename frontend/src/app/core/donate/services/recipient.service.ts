import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {environment} from '../../../../environments/environment';
import {map, retry} from 'rxjs/operators';
import {HttpClient, HttpParams} from '@angular/common/http';
import {RecipientModel} from '../models/recipient.model';

@Injectable({
  providedIn: 'root'
})
export class RecipientService {
  private tempRefObj = new RecipientModel();

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
    return this.http.get(environment.api_url_match + `recipients`, {params})
      .pipe(
        retry(2),
        map((response: any) => response)
      );
  }

  public getCanReceiveFrom(bloodType: number): Observable<[]> {
    return this.http.get<[]>(environment.api_url_match + `can-receive-from/${bloodType}`)
      .pipe(
        retry(2),
        map((response: any) => {
          return response.compatible_types
        })
      );
  }

  public getAllPublic(): Observable<RecipientModel[]> {
    return this.http.get<RecipientModel[]>(environment.api_url + 'recipient/public')
      .pipe(
        retry(2),
        map((response: any) => {
          const list = [];
          Object.keys(response.recipients).forEach((item, key) => {
            response.recipients[key].compatible = [1, 2, 3, 4, 5, 6, 7, 8];
            list.push(Object.assign(new RecipientModel(), response.recipients[key]))
          });
          return list;
        })
      );
  }

  public get(recipientId: number): Observable<RecipientModel> {
    return this.http.get<RecipientModel>(environment.api_url + 'recipient/' + recipientId)
      .pipe(
        retry(2),
        map((response: any) => Object.assign(new RecipientModel(), response[this.tempRefObj.getPrefix()]))
      );
  }

  public post(recipient: any): Observable<RecipientModel> {
    const body = {};
    body[this.tempRefObj.getPrefix()] = recipient;
    return this.http.post<RecipientModel>(environment.api_url + 'recipient', body)
      .pipe(
        map((response: any) => Object.assign(new RecipientModel(), response[this.tempRefObj.getPrefix()]))
      );
  }

  public put(recipient: RecipientModel): Observable<RecipientModel> {
    const body = {};
    body[this.tempRefObj.getPrefix()] = recipient;
    return this.http.put<RecipientModel>(environment.api_url + 'recipient/' + recipient.id, body)
      .pipe(
        map((response: any) => Object.assign(new RecipientModel(), response[this.tempRefObj.getPrefix()]))
      );
  }

  public delete(recipientId: number, answer: boolean, comment: string): Observable<any> {
    return this.http.patch<any>(environment.api_url + `recipient/${recipientId}/delete`,{
      answer,
      comment
    })
      .pipe(
        retry(2),
        map(response => recipientId)
      );
  }

  public getManagement() {
    return this.http.get<[]>(environment.api_url + 'recipient')
      .pipe(
        retry(2),
        map((response: any) => response.recipients)
      );
  }

  public verify(recipientId: number, isVerify: boolean): Observable<boolean> {
    return this.http.patch<boolean>(environment.api_url + `recipient/${recipientId}/verify`,
      {verified: isVerify})
      .pipe(
        retry(2),
        map((response: any) => response.ok)
      );
  }

  public setPublic(recipientId: number, isPublic: boolean): Observable<boolean> {
    return this.http.patch<boolean>(environment.api_url + `recipient/${recipientId}/public`,
      {public: isPublic})
      .pipe(
        retry(2),
        map((response: any) => response.ok)
      );
  }


  public setActivate(recipientId: number, isActivate: boolean): Observable<boolean> {
    return this.http.patch<boolean>(environment.api_url + `recipient/${recipientId}/activate`,
      {activate: isActivate})
      .pipe(
        retry(2),
        map((response: any) => response.ok)
      );
  }

}
