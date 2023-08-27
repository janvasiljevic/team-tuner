import { create } from 'zustand';
import { PagesSlice, createPagesSlices } from './pagesSlice';

// Apple sponsored store
export const useAppStore = create<PagesSlice>()((...a) => ({
  ...createPagesSlices(...a),
}));
