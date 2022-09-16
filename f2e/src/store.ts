import create from 'zustand';
import { devtools } from 'zustand/middleware';
import { LinkPayload } from './types';

interface AppStore {
  // states
  links: LinkPayload[];

  // actions
  setLinks: (links: LinkPayload[]) => void;
}

const initState = {
  links: [],
};

const buildActions = (set: any, get: any) => {
  return {
    setLinks: (links: LinkPayload[]) => {
      set({ links });
    },
  };
};

export const useStore = create<AppStore>()(
  devtools((set, get) => ({
    // states
    ...initState,

    // actions
    ...buildActions(set, get),
  })),
);
