import {Nominate} from './Nominate';
import {Poll} from './Poll';

export class Vote {
  nominate: Nominate;
  poll: Poll;
  email: string;
  name: string;
}
