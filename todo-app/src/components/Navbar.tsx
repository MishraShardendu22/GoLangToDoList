import React from 'react';
import { Github, Linkedin } from 'lucide-react';
import { motion, Variants } from 'framer-motion';

interface NavbarProps {
  className?: string;
}

const navVariants: Variants = {
  hidden: { opacity: 0, y: -20 },
  visible: { 
    opacity: 1, 
    y: 0,
    transition: { duration: 0.5 }
  }
};

const linkVariants: Variants = {
  hover: { 
    scale: 1.1,
    transition: { duration: 0.2 }
  }
};

const Navbar: React.FC<NavbarProps> = ({ className = '' }) => {
  return (
    <motion.div 
      initial="hidden"
      animate="visible"
      variants={navVariants}
      className={`w-full border-b bg-blue-400 border-border bg-background ${className}`}
    >
      <nav className="max-w-7xl mx-auto px-4 h-16 flex items-center justify-between">
        <motion.div 
          whileHover="hover"
          variants={linkVariants}
          className="flex items-center gap-2"
        >
          <a 
            href="https://www.linkedin.com/in/shardendumishra22/" 
            target="_blank" 
            rel="noopener noreferrer"
            className="flex items-center gap-2 text-foreground hover:text-primary transition-colors"
            aria-label="LinkedIn Profile"
          >
            <Linkedin className="w-5 h-5" />
            <span className="hidden sm:inline">LinkedIn</span>
          </a>
        </motion.div>

        <motion.h1 
          className="text-2xl font-bold text-primary"
          whileHover={{ scale: 1.05 }}
          transition={{ type: "spring", stiffness: 300 }}
        >
          Todo App
        </motion.h1>

        <div className="flex items-center gap-4">
          <motion.div 
            whileHover="hover"
            variants={linkVariants}
          >
            <a 
              href="https://github.com/ShardenduMishra22" 
              target="_blank" 
              rel="noopener noreferrer"
              className="flex items-center gap-2 text-foreground hover:text-primary transition-colors"
              aria-label="GitHub Profile"
            >
              <Github className="w-5 h-5" />
              <span className="hidden sm:inline">Github</span>
            </a>
          </motion.div>
        </div>
      </nav>
    </motion.div>
  );
};

export default Navbar;
