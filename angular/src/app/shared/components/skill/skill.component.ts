import { Component, Input } from '@angular/core';
import { Skill } from '../../../features/skills/models/skill.model';
import { NgForOf } from '@angular/common';


@Component({
  selector: 'app-skill',
  imports: [NgForOf],
  templateUrl: './skill.component.html',
  styleUrl: './skill.component.css',
  standalone: true,
})
export class SkillComponent {
  @Input() skill!: Skill;
}
