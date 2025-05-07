import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from '../../shared/components/navbar/navbar.component';
import { FooterComponent } from '../../shared/components/footer/footer.component';
import { SkillComponent } from '../../shared/components/skill/skill.component';


@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    NavbarComponent,
    FooterComponent,
    SkillComponent
  ]
})
export class PublicModule { }
