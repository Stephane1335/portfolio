import { Component, OnInit } from '@angular/core';
import { NavbarComponent } from '../../../shared/components/navbar/navbar.component';
import { FooterComponent } from '../../../shared/components/footer/footer.component';
import { AboutMe } from '../../../features/aboutme/models/aboutme.model';
import { AboutMeService } from '../../../features/aboutme/services/aboutme.service';
import { SkillComponent } from '../../../shared/components/skill/skill.component';
import { Skill } from '../../../features/skills/models/skill.model';
import { SkillService } from '../../../features/skills/services/skill.service';
import { NgForOf } from '@angular/common';


@Component({
  selector: 'app-home',
  imports: [NavbarComponent, FooterComponent, SkillComponent, NgForOf],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
  standalone: true,
})
export class HomeComponent implements OnInit {
 aboutMe: AboutMe = new AboutMe();
 skills: Skill[] =  [];
 constructor(private aboutMeService: AboutMeService, private skillService: SkillService){}
 ngOnInit(): void {
     this.getAboutMe();
     this.getSkill();
 }

 private getAboutMe(): void {
  this.aboutMeService.getAboutMes().subscribe({
    next: (response: AboutMe[]) => {
      if (response.length > 0) {
        this.aboutMe = new AboutMe(response[0]);
        console.log(this.aboutMe)
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
      console.log('Les Skills',this.skills)
    },
    error: (error: Error) => {
      console.error('Error of the data backup:',error)
    }
  })
}

}


