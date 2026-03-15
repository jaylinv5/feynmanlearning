import { request } from './api';
import type {
  KnowledgePoint,
  KnowledgePointDetail,
  KnowledgePointListItem,
  CreateKnowledgePointRequest,
  UpdateKnowledgePointRequest,
  KnowledgeQueryParams,
} from '@/types/knowledge';

interface PageResult<T> {
  total: number;
  page: number;
  pageSize: number;
  data: T[];
}

// 知识点API
export const knowledgeApi = {
  // 创建知识点
  create: (data: CreateKnowledgePointRequest) => {
    return request.post<KnowledgePoint>('/knowledge/create', data);
  },

  // 更新知识点
  update: (data: UpdateKnowledgePointRequest) => {
    return request.put<KnowledgePoint>('/knowledge/update', data);
  },

  // 删除知识点
  delete: (id: number) => {
    return request.delete(`/knowledge/delete/${id}`);
  },

  // 获取知识点详情
  getDetail: (id: number) => {
    return request.get<KnowledgePointDetail>(`/knowledge/detail/${id}`);
  },

  // 分页查询知识点列表
  list: (params: KnowledgeQueryParams) => {
    return request.get<PageResult<KnowledgePoint>>('/knowledge/list', { params });
  },

  // 根据学科和年级获取知识点列表
  getBySubjectAndGrade: (subject: string, grade: number) => {
    return request.get<KnowledgePointListItem[]>(`/knowledge/subject/${subject}/grade/${grade}`);
  },

  // 批量更新知识点状态
  batchUpdateStatus: (ids: number[], status: number) => {
    return request.put('/knowledge/batch/status', { ids, status });
  },

  // 批量导入知识点
  batchImport: (data: CreateKnowledgePointRequest[]) => {
    return request.post<{ import_count: number }>('/knowledge/batch/import', data);
  },
};
