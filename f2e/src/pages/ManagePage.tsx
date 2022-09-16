import { useEffect } from 'react';
import { Shortener, UrlTable } from '../components';
import { fetchLinks } from '../api/links';
import { useStore } from '../store';

const ManagePage = () => {
  const [links, setLinks] = useStore((state) => [state.links, state.setLinks]);

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
