import { useEffect } from 'react';
import { Shortener, UrlTable } from '../components';
import { fetchLinks } from '../api/links';
import { useStore } from '../store';

const ManagePage = () => {
  const links = useStore((state) => state.links);
  const setLinks = useStore((state) => state.setLinks);

  useEffect(() => {
    fetchLinks().then((payload) => {
      setLinks(payload);
    });
  }, [setLinks]);

  return (
    <div>
      <Shortener />
      <UrlTable links={links} />
    </div>
  );
};

export default ManagePage;
