import {Nominate} from './Nominate';

export class Poll {
  id: number;
  year: number;
  duration: number;
  nominates: Nominate[];
}
