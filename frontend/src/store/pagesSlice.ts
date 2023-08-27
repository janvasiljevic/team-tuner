import { StateCreator } from 'zustand';
import { OutCourseOut } from '../api/model';

export interface PagesSlice {
  selectedCourse: OutCourseOut | null;
  setSelectedCourse: (course: OutCourseOut | null) => void;
}

export const createPagesSlices: StateCreator<PagesSlice, [], [], PagesSlice> = (
  set,
) => ({
  selectedCourse: null,
  setSelectedCourse(course) {
    set((state) => ({ ...state, selectedCourse: course }));
  },
});
