export interface Workspace {
  id: number;
  name: string;
  notification_type: number;
  notification_id: number | null;
  description: string | null;
  created_at: string;
  updated_at: string;
}
