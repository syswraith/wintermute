import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

type Link = {
  ID: number;
  ShortURL: string;
  LongURL: string;
};

export function Dashboard() {
  const [links, setLinks] = useState<Link[]>([]);

  useEffect(() => {
    const fetchJsonLinks = async () => {
      const url = "http://localhost:31337/dashboard";
      const res = await fetch(url);
      const jsonLinks = await res.json();
      setLinks(jsonLinks);
    };

    fetchJsonLinks();
  }, []);

  return (
    <div className="flex justify-center mt-10 px-8">
      <div className="w-full max-w-4xl">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="text-xl">Short URL</TableHead>
              <TableHead className="text-xl">Long URL</TableHead>
              <TableHead className="text-xl text-right">Expiry</TableHead>
            </TableRow>
          </TableHeader>

          <TableBody>
            {links.map((link) => (
              <TableRow key={link.ID}>
                <TableCell className="font-mono">
                  <a href="">http://localhost:31337/{link.ShortURL}</a>
                </TableCell>
                <TableCell className="truncate max-w-sm">
                  <a href="">{link.LongURL}</a>
                </TableCell>
                <TableCell className="text-right">-</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
