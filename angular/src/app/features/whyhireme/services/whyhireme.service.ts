import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { WhyHireMe } from '../models/whyhireme.model';
import { environment } from '../../../../environnements/environnement';

@Injectable({
  providedIn: 'root'
})
export class WhyHireMeService {
  private apiUrl = `${environment.apiUrl}/whyhireme`;

  constructor(private http: HttpClient) {}

  getWhyHireMes(): Observable<WhyHireMe[]> {
    return this.http.get<WhyHireMe[]>(this.apiUrl).pipe(
      catchError(this.handleError)
    );
  }

  getWhyHireMe(id: string): Observable<WhyHireMe> {
    return this.http.get<WhyHireMe>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  createWhyHireMe(WhyHireMe: WhyHireMe): Observable<WhyHireMe> {
    return this.http.post<WhyHireMe>(this.apiUrl, WhyHireMe).pipe(
      catchError(this.handleError)
    );
  }

  updateWhyHireMe(id: string, WhyHireMe: WhyHireMe): Observable<WhyHireMe> {
    return this.http.put<WhyHireMe>(`${this.apiUrl}/${id}`, WhyHireMe).pipe(
      catchError(this.handleError)
    );
  }

  deleteWhyHireMe(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    // Gérer l'erreur de manière appropriée
    return throwError(() => new Error('Une erreur est survenue.'));
  }
}
