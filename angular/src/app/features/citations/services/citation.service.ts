import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Citation } from '../models/citation.model';
import { environment } from '../../../../environnements/environnement';

@Injectable({
  providedIn: 'root'
})
export class CitationService {
  private apiUrl = `${environment.apiUrl}/citations`;

  constructor(private http: HttpClient) {}

  getCitations(): Observable<Citation[]> {
    return this.http.get<Citation[]>(this.apiUrl).pipe(
      catchError(this.handleError)
    );
  }

  getCitation(id: string): Observable<Citation> {
    return this.http.get<Citation>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  createCitation(Citation: Citation): Observable<Citation> {
    return this.http.post<Citation>(this.apiUrl, Citation).pipe(
      catchError(this.handleError)
    );
  }

  updateCitation(id: string, Citation: Citation): Observable<Citation> {
    return this.http.put<Citation>(`${this.apiUrl}/${id}`, Citation).pipe(
      catchError(this.handleError)
    );
  }

  deleteCitation(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    // Gérer l'erreur de manière appropriée
    return throwError(() => new Error('Une erreur est survenue.'));
  }
}
