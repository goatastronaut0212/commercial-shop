{ config, pkgs, ... }:

{
  programs.neovim = {
    enable = true;
    vimAlias = true;
    extraConfig = ''
      luafile $NVIM_DIR/settings.lua
      luafile $NVIM_DIR/bufferline.lua
      luafile $NVIM_DIR/lualine.lua
      luafile $NVIM_DIR/nvim-tree.lua
      set background=light
      colorscheme PaperColor
      NvimTreeToggle
    '';
    plugins = with pkgs.vimPlugins; [
      # plugins written in VimScript 
      papercolor-theme
      vim-nix

      # plugins written in Lua
      bufferline-nvim
      lualine-nvim
      nvim-tree-lua
      nvim-web-devicons
    ];
  };
}
