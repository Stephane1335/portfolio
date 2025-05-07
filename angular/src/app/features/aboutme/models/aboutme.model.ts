// src/app/features/aboutme/models/aboutme.model.ts
export class AboutMe {
  _id?: string;
  quote?: string;
  description?: string;
  position?: string;
  name?: string;
  signature?: string;

  constructor(init?: Partial<AboutMe>) {
    Object.assign(this, init);
  }
}
