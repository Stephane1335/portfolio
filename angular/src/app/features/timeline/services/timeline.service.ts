import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Timeline } from '../models/timeline.model';
import { environment } from '../../../../environnements/environnement';

@Injectable({
  providedIn: 'root'
})
export class TimelineService {
  private apiUrl = `${environment.apiUrl}/timelines`;

  constructor(private http: HttpClient) {}

  getTimelines(): Observable<Timeline[]> {
    return this.http.get<Timeline[]>(this.apiUrl).pipe(
      catchError(this.handleError)
    );
  }

  getTimeline(id: string): Observable<Timeline> {
    return this.http.get<Timeline>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  createTimeline(Timeline: Timeline): Observable<Timeline> {
    return this.http.post<Timeline>(this.apiUrl, Timeline).pipe(
      catchError(this.handleError)
    );
  }

  updateTimeline(id: string, Timeline: Timeline): Observable<Timeline> {
    return this.http.put<Timeline>(`${this.apiUrl}/${id}`, Timeline).pipe(
      catchError(this.handleError)
    );
  }

  deleteTimeline(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    // Gérer l'erreur de manière appropriée
    return throwError(() => new Error('Une erreur est survenue.'));
  }
}
