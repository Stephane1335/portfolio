import { Component, Input } from '@angular/core';
import { WhyHireMe } from '../../../features/whyhireme/models/whyhireme.model';

@Component({
  selector: 'app-whyhireme',
  imports: [],
  templateUrl: './whyhireme.component.html',
  styleUrl: './whyhireme.component.css'
})
export class WhyhiremeComponent {
@Input() whyhireme!: WhyHireMe
}
