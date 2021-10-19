export interface GoalDto {
  id: string;
  order_id: number;
  text: string;
  color: string;
}

export interface MoveGoalDto {
  goalId: string;
  afterGoalId: string | undefined;
}

export interface CreateGoalDto {
  text: string;
}
