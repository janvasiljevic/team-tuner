/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * TT API
 * Team tuner API documentation
 * OpenAPI spec version: 1.0
 */
import type { GetStudentSortField } from './getStudentSortField';
import type { GetStudentSortOrder } from './getStudentSortOrder';

export type GetStudentParams = {
  completedQuestioner?: boolean;
  courseId?: string;
  page?: number;
  pageSize?: number;
  sortField?: GetStudentSortField;
  sortOrder?: GetStudentSortOrder;
};
