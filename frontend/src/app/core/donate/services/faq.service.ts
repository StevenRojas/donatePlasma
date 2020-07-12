import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { retry, map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class FaqService {

  constructor(
    private http: HttpClient
  ) { }

  public comments(email: string, comment: string): Observable<any> {
    return this.http.post<any>(environment.api_url_simple + '/api/comments', {
      email, comment
    })
    .pipe(
      retry(2),
      map( response => response)
    );
  }

  public recruit(email: string, comment: string): Observable<any> {
    return this.http.post<any>(environment.api_url_simple + '/api/recruit', {
      email, comment
    })
    .pipe(
      retry(2),
      map( response => response)
    );
  }
}
