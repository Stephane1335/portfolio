import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Skill } from '../models/skill.model';
import { environment } from '../../../../environnements/environnement';

@Injectable({
  providedIn: 'root'
})
export class SkillService {
  private apiUrl = `${environment.apiUrl}/skills`;

  constructor(private http: HttpClient) {}

  getSkills(): Observable<Skill[]> {
    return this.http.get<Skill[]>(this.apiUrl).pipe(
      catchError(this.handleError)
    );
  }

  getSkill(id: string): Observable<Skill> {
    return this.http.get<Skill>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  createSkill(Skill: Skill): Observable<Skill> {
    return this.http.post<Skill>(this.apiUrl, Skill).pipe(
      catchError(this.handleError)
    );
  }

  updateSkill(id: string, Skill: Skill): Observable<Skill> {
    return this.http.put<Skill>(`${this.apiUrl}/${id}`, Skill).pipe(
      catchError(this.handleError)
    );
  }

  deleteSkill(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    // Gérer l'erreur de manière appropriée
    return throwError(() => new Error('Une erreur est survenue.'));
  }
}
