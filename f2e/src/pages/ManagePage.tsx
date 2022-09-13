import { useEffect, useState } from 'react';
import { Shortener, UrlTable } from '../components';
import { getLinks, LinkPayload } from '../api/links';

const ManagePage = () => {
  const [links, setLinks] = useState<LinkPayload[]>([]);

  useEffect(() => {
    getLinks().then((payload) => {
      setLinks(payload);
    });
  }, []);

  return (
    <div>
      <h1>It is manage page!</h1>
      <Shortener />
      <UrlTable links={links} />
    </div>
  );
};

export default ManagePage;
