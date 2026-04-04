import { useEffect, useState } from "react";
import { Navbar } from "@/components/Navbar";
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
  Expiry: string | null;
};

export function Dashboard() {
  const [links, setLinks] = useState<Link[]>([]);
  const [now, setNow] = useState(Date.now());

  useEffect(() => {
    const fetchJsonLinks = async () => {
      const url = "http://localhost:31337/dashboard";
      const res = await fetch(url, {
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("token")}`,
        },
      });
      if (res.ok) {
        const jsonLinks = await res.json();
        setLinks(jsonLinks);
      }
    };

    fetchJsonLinks();
  }, []);

  useEffect(() => {
    const timer = setInterval(() => {
      setNow(Date.now());
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  const formatCountdown = (expiry: string | null) => {
    if (!expiry) return "Never";

    const diff = new Date(expiry).getTime() - Date.now();

    if (diff <= 0) return "Expired";

    const seconds = Math.floor(diff / 1000) % 60;
    const minutes = Math.floor(diff / (1000 * 60)) % 60;
    const hours = Math.floor(diff / (1000 * 60 * 60)) % 24;
    const days = Math.floor(diff / (1000 * 60 * 60 * 24));

    return `${days}d ${hours}h ${minutes}m ${seconds}s`;
  };

  return (
    <div className="flex flex-col min-h-screen">
      <Navbar />
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
                  <a href={`http://localhost:31337/${link.ShortURL}`}>
                    http://localhost:31337/{link.ShortURL}
                  </a>
                </TableCell>

                <TableCell className="truncate max-w-sm">
                  <a href={link.LongURL}>{link.LongURL}</a>
                </TableCell>

                <TableCell className="text-right">
                  {formatCountdown(link.Expiry)}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
    </div>
  );
}
