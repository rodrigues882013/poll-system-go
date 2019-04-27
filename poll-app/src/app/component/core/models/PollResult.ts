import {Nominate} from './Nominate';

class ResultArray {
  nominate: Nominate;
  votes: number;
}

export class PollResult {
  pollId: number;
  results: ResultArray[];
}
