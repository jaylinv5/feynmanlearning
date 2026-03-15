// 知识点类型定义
export interface Example {
  id: string;
  title: string;
  content: string;
  analysis: string;
  answer: string;
  difficulty: number;
}

export interface Exercise {
  id: string;
  type: number; // 1-单选,2-多选,3-判断,4-填空,5-简答
  question: string;
  options?: string[];
  answer: string;
  analysis: string;
  difficulty: number;
  score: number;
}

export interface FeynmanGuide {
  id: string;
  question: string;
  difficulty: number;
  keyPoint: string;
}

export interface KnowledgePoint {
  id: number;
  subject: string;
  grade: number;
  chapter: string;
  chapterOrder: number;
  name: string;
  code: string;
  difficulty: number;
  estimatedTime: number;
  content: string;
  status: number;
  tags: string;
  createdAt: string;
  updatedAt: string;
}

export interface KnowledgePointDetail extends KnowledgePoint {
  examples: Example[];
  exercises: Exercise[];
  feynmanGuide: FeynmanGuide[];
  preRequires: number[];
}

export interface KnowledgePointListItem extends KnowledgePoint {
  learned: boolean;
  mastered: boolean;
}

export interface CreateKnowledgePointRequest {
  subject: string;
  grade: number;
  chapter: string;
  chapterOrder: number;
  name: string;
  difficulty: number;
  estimatedTime: number;
  content: string;
  examples: Example[];
  exercises: Exercise[];
  feynmanGuide: FeynmanGuide[];
  preRequires: number[];
  tags: string;
}

export interface UpdateKnowledgePointRequest extends Partial<CreateKnowledgePointRequest> {
  id: number;
  status?: number;
}

export interface KnowledgeQueryParams {
  page?: number;
  pageSize?: number;
  subject?: string;
  grade?: number;
  status?: number;
}
