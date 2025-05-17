import { Component, OnInit } from '@angular/core';
import { NavbarComponent } from '../../../shared/components/navbar/navbar.component';
import { FooterComponent } from '../../../shared/components/footer/footer.component';
import { AboutMe } from '../../../features/aboutme/models/aboutme.model';
import { AboutMeService } from '../../../features/aboutme/services/aboutme.service';
import { SkillComponent } from '../../../shared/components/skill/skill.component';
import { Skill } from '../../../features/skills/models/skill.model';
import { SkillService } from '../../../features/skills/services/skill.service';
import { NgForOf, SlicePipe } from '@angular/common';
import { TimelineComponent } from '../../../shared/components/timeline/timeline.component';
import { WhyHireMeService } from '../../../features/whyhireme/services/whyhireme.service';
import { WhyHireMe } from '../../../features/whyhireme/models/whyhireme.model';
import { Project } from '../../../features/projects/models/project.model';
import { ProjectService } from '../../../features/projects/services/project.service';



@Component({
  selector: 'app-home',
  imports: [FooterComponent, SkillComponent, TimelineComponent, NgForOf, SlicePipe],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
  standalone: true,
})
export class HomeComponent implements OnInit {
 aboutMe: AboutMe = new AboutMe();
 skills: Skill[] =  [];
 whyhiremes: WhyHireMe[] = [];
 projects: Project[] = [];
 constructor(private aboutMeService: AboutMeService, 
             private skillService: SkillService,
             private whyhiremeservice: WhyHireMeService,
             private projectService: ProjectService,
             ){}
 ngOnInit(): void {
     this.getAboutMe();
     this.getSkill();
     this.getWhyhireme();
     this.getProject(); 
 }

 private getAboutMe(): void {
  this.aboutMeService.getAboutMes().subscribe({
    next: (response: AboutMe[]) => {
      if (response.length > 0) {
        this.aboutMe = new AboutMe(response[0]);
      }
    },
    error: (error: Error) => {
      console.error('Erreur lors de la récupération des données:', error);
    }
  });
}

private getSkill(): void {
  this.skillService.getSkills().subscribe({
    next: (response: any) => {
      this.skills = response.data
    },
    error: (error: Error) => {
      console.error('Error of the data backup:',error)
    }
  })
}

private getWhyhireme(): void {
  this.whyhiremeservice.getWhyHireMes().subscribe({
    next:(response: any) =>{
      this.whyhiremes = response.data
    },
    error: (error: Error) => {
      console.error('Error of the data backup:',error)
    }
  })
}

private getProject(): void {
  this.projectService.getProjects().subscribe({
    next:(response: any) =>{
      this.projects = response.data
    },
    error: (error: Error) => {
      console.error('Error of the data backup:',error)
    }
  })
}

}


