import { Component, OnInit } from '@angular/core';
import { NavbarComponent } from '../../../shared/components/navbar/navbar.component';
import { FooterComponent } from '../../../shared/components/footer/footer.component';
import { AboutMe } from '../../../features/aboutme/models/aboutme.model';
import { AboutMeService } from '../../../features/aboutme/services/aboutme.service';

@Component({
  selector: 'app-home',
  imports: [NavbarComponent, FooterComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
 aboutMe: AboutMe = new AboutMe();
 constructor(private aboutMeService: AboutMeService){}
 ngOnInit(): void {
     this.getAboutMe();
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

}


