import { z } from 'zod';

const initDNSServersEvent = z.object({
  dns_servers: z.array(
    z.object({
      ip_and_port: z.string(),
      hostname: z.string(),
      protocol: z.number(),
      responds_a: z.union([z.boolean(), z.null()]),
      responds_aaaa: z.union([z.boolean(), z.null()]),
    })
  ),
});

export type initDNSServersEventType = z.infer<typeof initDNSServersEvent>;

export const parseInitDNSServersEventType = (
  unk: unknown
): initDNSServersEventType => {
  return initDNSServersEvent.parse(unk);
};
