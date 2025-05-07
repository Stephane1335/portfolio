import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Project } from '../models/project.model';
import { environment } from '../../../../environnements/environnement';

@Injectable({
  providedIn: 'root'
})
export class ProjectService {
  private apiUrl = `${environment.apiUrl}/projects`;

  constructor(private http: HttpClient) {}

  getProjects(): Observable<Project[]> {
    return this.http.get<Project[]>(this.apiUrl).pipe(
      catchError(this.handleError)
    );
  }

  getProject(id: string): Observable<Project> {
    return this.http.get<Project>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  createProject(project: Project): Observable<Project> {
    return this.http.post<Project>(this.apiUrl, project).pipe(
      catchError(this.handleError)
    );
  }

  updateProject(id: string, project: Project): Observable<Project> {
    return this.http.put<Project>(`${this.apiUrl}/${id}`, project).pipe(
      catchError(this.handleError)
    );
  }

  deleteProject(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: HttpErrorResponse) {
    // Gérer l'erreur de manière appropriée
    return throwError(() => new Error('Une erreur est survenue.'));
  }
}
