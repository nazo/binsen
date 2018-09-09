export interface Post {
  id: number;
  workspace_id: number;
  title: string;
  body: string;
  creator_user_id: number;
  editor_user_id: number;
  edit_status: number;
  created_at: string;
  updated_at: string;
}