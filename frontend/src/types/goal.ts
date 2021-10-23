export interface GoalDto {
  id: string;
  text: string;
  rank: number;
}

export interface MoveGoalDto {
  goalId: string;
  afterGoalId: string | undefined;
}

export interface CreateGoalDto {
  text: string;
}

export interface ModifyGoalDto {
  text: string;
}
