import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { AboutMe } from '../models/aboutme.model';
import { environment } from '../../../../environnements/environnement';

@Injectable({
  providedIn: 'root'
})
export class AboutMeService {
  private apiUrl = `${environment.apiUrl}/aboutme`;

  constructor(private http: HttpClient) {}

  getAboutMes(): Observable<AboutMe[]> {
    return this.http.get<{ data: AboutMe[] }>(this.apiUrl).pipe(
      map(response => response.data),
      catchError(this.handleError)
    );
  }


  getAboutMe(id: string): Observable<AboutMe> {
    return this.http.get<{ data: AboutMe[] }>(`${this.apiUrl}/${id}`).pipe(
      map(response => response.data[0]),
      catchError(this.handleError)
    );
  }

  createAboutMe(aboutMe: AboutMe): Observable<AboutMe> {
    return this.http.post<{ data: AboutMe }>(this.apiUrl, aboutMe).pipe(
      map(response => response.data),
      catchError(this.handleError)
    );
  }

  updateAboutMe(id: string, aboutMe: AboutMe): Observable<AboutMe> {
    return this.http.put<{ data: AboutMe }>(`${this.apiUrl}/${id}`, aboutMe).pipe(
      map(response => response.data),
      catchError(this.handleError)
    );
  }

  deleteAboutMe(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    console.error('Erreur HTTP:', {
      status: error.status,
      message: error.message,
      error: error.error
    });

    let errorMessage = 'Une erreur est survenue.';

    if (error.error instanceof ErrorEvent) {
      // Erreur côté client
      errorMessage = `Erreur côté client: ${error.error.message}`;
    } else {
      // Erreur côté serveur
      errorMessage = `Erreur côté serveur: Code ${error.status}, Message: ${error.message}`;
    }

    return throwError(() => new Error(errorMessage));
  }
}
