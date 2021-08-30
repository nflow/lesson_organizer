export interface ContentDto {
  id: string;
  boardId: string;
  methodId: string;
  text: string;
  order: number;
}

export interface CreateContentDto {
  text: string;
}
