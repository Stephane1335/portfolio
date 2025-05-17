import { Component, Input } from '@angular/core';
import { Project } from '../../../features/projects/models/project.model';

@Component({
  selector: 'app-project',
  imports: [],
  templateUrl: './project.component.html',
  styleUrl: './project.component.css'
})
export class ProjectComponent {
@Input() project!: Project
}
