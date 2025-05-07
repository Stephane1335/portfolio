export class Skill {
    id?: string;
    title?: string;
    techno?: string[];
    description?: string;
    level?: string;
    color?: string;

    constructor(init?: Partial<Skill>) {
      Object.assign(this, init);
    }
  }
  